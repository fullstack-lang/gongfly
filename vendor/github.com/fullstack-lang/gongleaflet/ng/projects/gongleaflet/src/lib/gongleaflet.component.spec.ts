import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongleafletComponent } from './gongleaflet.component';

describe('GongleafletComponent', () => {
  let component: GongleafletComponent;
  let fixture: ComponentFixture<GongleafletComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GongleafletComponent]
    });
    fixture = TestBed.createComponent(GongleafletComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
