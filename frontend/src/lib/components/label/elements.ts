import { get, writable, type Writable } from 'svelte/store';

export interface Element {
	z: number;
	x: number;
	y: number;
	height: number;
	width: number;
	draggable?: boolean;
	attributes?: Record<string, Attribute<'color' | 'text'>>;

	draw(ctx: CanvasRenderingContext2D): void;
	click?(): void;
	resize?(by: number): void;
}

type AttributeType = 'text' | 'color';

interface Attribute<T extends AttributeType> {
	label: string;
	type: T;
	value: string;
	hidden?: boolean;
}

export const moveInBounds = (
	element: Element,
	x: number,
	y: number,
	width: number,
	height: number
) => {
	if (element.x < x) {
		element.x = x;
	}
	if (element.y < y) {
		element.y = y;
	}
	if (element.x + element.width > x + width) {
		element.x = x + width - element.width;
	}
	if (element.y + element.height > y + height) {
		element.y = y + height - element.height;
	}
};

export class Rectangle implements Element {
	z: number;
	x: number;
	y: number;
	height: number;
	width: number;
	draggable: boolean;

	constructor(x: number, y: number, height: number, width: number, draggable = false, z = 0) {
		this.x = x;
		this.y = y;
		this.height = height;
		this.width = width;
		this.draggable = draggable;
		this.z = z;
	}

	draw(ctx: CanvasRenderingContext2D) {
		ctx.beginPath();
		ctx.fillRect(this.x, this.y, this.width, this.height);
	}

	click() {
		console.log('click');
	}

	resize(by: number) {
		this.height += by * (this.height / this.width);
		this.width += by * (this.width / this.height);
	}
}

interface TextAttributes {
	[key: string]: Attribute<'color' | 'text'>;
	text: Attribute<'text'>;
	size: Attribute<'text'>;
	font: Attribute<'text'>;
	color: Attribute<'color'>;
}

export class Text implements Element {
	z: number;
	x: number;
	y: number;
	height: number = 0;
	width: number = 0;
	rotation: number;
	draggable: boolean;
	attributes: TextAttributes;

	constructor(
		x: number,
		y: number,
		attributes: TextAttributes,
		draggable = false,
		z = 0,
		rotation = 0
	) {
		this.x = x;
		this.y = y;
		this.attributes = attributes;
		this.draggable = draggable;
		this.z = z;
		this.rotation = rotation;
	}

	draw(ctx: CanvasRenderingContext2D): void {
		ctx.font = `${this.attributes.size.value} ${this.attributes.font.value}`;
		ctx.fillStyle = this.attributes.color.value;

		const textMetrics = ctx.measureText(this.attributes.text.value);
		this.width = textMetrics.width;
		this.height = textMetrics.actualBoundingBoxAscent + textMetrics.actualBoundingBoxDescent;

		ctx.fillText(this.attributes.text.value, this.x, this.y + this.height);
	}
}

interface ImageAttributes {
	[key: string]: Attribute<'text'>;
	src: Attribute<'text'>;
}

export class ImageElement implements Element {
	z: number;
	x: number;
	y: number;
	height: number;
	width: number;
	draggable: boolean;
	attributes: ImageAttributes;

	loaded = false;
	image: HTMLImageElement;

	constructor(
		x: number,
		y: number,
		attributes: ImageAttributes,
		width: number = 0,
		height: number = 0,
		draggable = false,
		z = 0
	) {
		this.x = x;
		this.y = y;
		this.height = height;
		this.width = width;
		this.attributes = attributes;
		this.draggable = draggable;
		this.z = z;

		this.image = new Image();
		this.image.src = this.attributes.src.value;
		this.image.onload = () => {
			if (this.width === 0) {
				this.width = this.image.width;
			}
			if (this.height === 0) {
				this.height = this.image.height;
			}
			this.loaded = true;
		};
	}

	draw(ctx: CanvasRenderingContext2D): void {
		if (this.loaded) {
			ctx.drawImage(this.image, this.x, this.y, this.width, this.height);
		} else {
			ctx.fillStyle = 'grey';
			ctx.fillRect(this.x, this.y, this.width, this.height);
		}
	}

	resize(by: number): void {
		this.height += by * (this.height / this.width);
		this.width += by * (this.width / this.height);
	}
}

// Centers an element in the given width and height
export const centerElement = (element: Element, width: number, height: number) => {
	element.x = (width - element.width) / 2;
	element.y = (height - element.height) / 2;
};

export class Container {
	elements: Element[] = [];
	focused: Writable<Element | null> = writable(null);

	constructor() {}

	push(element: Element) {
		this.elements.push(element);
	}

	clear() {
		this.elements = [];
	}

	atPoint(x: number, y: number): Element[] {
		return this.elements.filter((e) => {
			return e.x <= x && e.x + e.width >= x && e.y <= y && e.y + e.height >= y;
		});
	}

	singleAtPoint(x: number, y: number): Element | null {
		const elements = this.atPoint(x, y);
		return elements.length > 0 ? elements[elements.length - 1] : null;
	}

	sortByZ() {
		this.elements.sort((a, b) => a.z - b.z);
	}

	draw(ctx: CanvasRenderingContext2D) {
		this.sortByZ();
		this.elements.forEach((e) => {
			e.draw(ctx);
			moveInBounds(e, 0, 0, ctx.canvas.width, ctx.canvas.height);
		});

		const focused = get(this.focused);
		if (focused) {
			ctx.beginPath();
			ctx.strokeStyle = 'red';
			ctx.lineWidth = 4;
			ctx.rect(focused.x - 2, focused.y - 2, focused.width + 4, focused.height + 4);
			ctx.stroke();
		}
	}

	click(x: number, y: number) {
		const element = this.singleAtPoint(x, y);
		if (element) {
			element.click && element.click();
			this.focused.set(element);
		} else {
			this.focused.set(null);
		}
	}

	hover(x: number, y: number) {
		const element = this.singleAtPoint(x, y);
		if (element?.draggable) {
			document.body.style.cursor = 'move';
		} else if (element?.resize) {
			document.body.style.cursor = 'nwse-resize';
		} else {
			document.body.style.cursor = 'default';
		}
	}

	drag(x: number, y: number, movementX: number, movementY: number) {
		const element = this.singleAtPoint(x, y);
		if (element && element.draggable) {
			element.x += movementX;
			element.y += movementY;
		}
	}

	resize(x: number, y: number, by: number) {
		const element = this.singleAtPoint(x, y);
		if (element && element.resize) {
			element.resize(by);
		}
	}
}
