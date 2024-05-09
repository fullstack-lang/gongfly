import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongflyComponent } from './gongfly.component';

describe('GongflyComponent', () => {
  let component: GongflyComponent;
  let fixture: ComponentFixture<GongflyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongflyComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongflyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
