// generated from ng_file_enum.ts.go
export enum ElementType {
	// insertion point	
	PARAGRAPH = "Paragraph",
	TITLE = "Title",
	TABLE = "Table",
}

export interface ElementTypeSelect {
	value: string;
	viewValue: string;
}

export const ElementTypeList: ElementTypeSelect[] = [ // insertion point	
	{ value: ElementType.PARAGRAPH, viewValue: "Paragraph" },
	{ value: ElementType.TITLE, viewValue: "Title" },
	{ value: ElementType.TABLE, viewValue: "Table" },
];
