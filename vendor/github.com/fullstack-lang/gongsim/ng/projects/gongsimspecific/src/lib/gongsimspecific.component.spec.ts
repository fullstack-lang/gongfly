import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsimspecificComponent } from './gongsimspecific.component';

describe('GongsimspecificComponent', () => {
  let component: GongsimspecificComponent;
  let fixture: ComponentFixture<GongsimspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongsimspecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongsimspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
