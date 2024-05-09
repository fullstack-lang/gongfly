import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongflyDashboardComponent } from './gongfly-dashboard.component';

describe('GongflyDashboardComponent', () => {
  let component: GongflyDashboardComponent;
  let fixture: ComponentFixture<GongflyDashboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongflyDashboardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongflyDashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
