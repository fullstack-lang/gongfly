import { TestBed } from '@angular/core/testing';

import { GongsimspecificService } from './gongsimspecific.service';

describe('GongsimspecificService', () => {
  let service: GongsimspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongsimspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
