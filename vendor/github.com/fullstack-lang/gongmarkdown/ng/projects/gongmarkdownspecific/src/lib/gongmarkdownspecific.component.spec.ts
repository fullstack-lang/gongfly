import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmarkdownspecificComponent } from './gongmarkdownspecific.component';

describe('GongmarkdownspecificComponent', () => {
  let component: GongmarkdownspecificComponent;
  let fixture: ComponentFixture<GongmarkdownspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongmarkdownspecificComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GongmarkdownspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
