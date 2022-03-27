import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmarkdownspecificMarkdownComponent } from './gongmarkdownspecific-markdown.component';

describe('GongmarkdownspecificMarkdownComponent', () => {
  let component: GongmarkdownspecificMarkdownComponent;
  let fixture: ComponentFixture<GongmarkdownspecificMarkdownComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongmarkdownspecificMarkdownComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GongmarkdownspecificMarkdownComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
