import { TestBed } from '@angular/core/testing';

import { GongflyService } from './gongfly.service';

describe('GongflyService', () => {
  let service: GongflyService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongflyService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
