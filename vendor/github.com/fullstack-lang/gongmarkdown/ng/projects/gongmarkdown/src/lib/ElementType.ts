// generated from ng_file_enum.ts.go
export enum ElementType {
	// insertion point	
	PARAGRAPH = "Paragraph",
	TABLE = "Table",
	TITLE = "Title",
}

export interface ElementTypeSelect {
	value: string;
	viewValue: string;
}

export const ElementTypeList: ElementTypeSelect[] = [ // insertion point	
	{ value: ElementType.PARAGRAPH, viewValue: "Paragraph" },
	{ value: ElementType.TABLE, viewValue: "Table" },
	{ value: ElementType.TITLE, viewValue: "Title" },
];
