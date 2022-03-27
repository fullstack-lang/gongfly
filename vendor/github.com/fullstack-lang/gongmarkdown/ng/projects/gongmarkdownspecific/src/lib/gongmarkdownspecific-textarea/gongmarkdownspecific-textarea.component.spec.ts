import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmarkdownspecificTextareaComponent } from './gongmarkdownspecific-textarea.component';

describe('GongmarkdownspecificTextareaComponent', () => {
  let component: GongmarkdownspecificTextareaComponent;
  let fixture: ComponentFixture<GongmarkdownspecificTextareaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongmarkdownspecificTextareaComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GongmarkdownspecificTextareaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
