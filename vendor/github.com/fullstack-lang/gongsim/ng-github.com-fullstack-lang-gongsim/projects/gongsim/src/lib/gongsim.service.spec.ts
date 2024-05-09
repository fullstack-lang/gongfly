import { TestBed } from '@angular/core/testing';

import { GongsimService } from './gongsim.service';

describe('GongsimService', () => {
  let service: GongsimService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsimService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
