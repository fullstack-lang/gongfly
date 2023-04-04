import { TestBed } from '@angular/core/testing';

import { GongflyspecificService } from './gongflyspecific.service';

describe('GongflyspecificService', () => {
  let service: GongflyspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongflyspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
