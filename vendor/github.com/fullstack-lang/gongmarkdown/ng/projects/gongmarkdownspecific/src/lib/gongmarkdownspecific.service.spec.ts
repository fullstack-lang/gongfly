import { TestBed } from '@angular/core/testing';

import { GongmarkdownspecificService } from './gongmarkdownspecific.service';

describe('GongmarkdownspecificService', () => {
  let service: GongmarkdownspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongmarkdownspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
