import { TestBed } from '@angular/core/testing';

import { GongleafletService } from './gongleaflet.service';

describe('GongleafletService', () => {
  let service: GongleafletService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongleafletService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
