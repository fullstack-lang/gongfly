import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongsimComponent } from './gongsim.component';

describe('GongsimComponent', () => {
  let component: GongsimComponent;
  let fixture: ComponentFixture<GongsimComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongsimComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongsimComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
