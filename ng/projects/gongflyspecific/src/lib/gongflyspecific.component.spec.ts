import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongflyspecificComponent } from './gongflyspecific.component';

describe('GongflyspecificComponent', () => {
  let component: GongflyspecificComponent;
  let fixture: ComponentFixture<GongflyspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongflyspecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongflyspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
