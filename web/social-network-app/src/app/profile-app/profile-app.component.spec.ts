import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProfileAppComponent } from './profile-app.component';

describe('ProfileAppComponent', () => {
  let component: ProfileAppComponent;
  let fixture: ComponentFixture<ProfileAppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProfileAppComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProfileAppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
