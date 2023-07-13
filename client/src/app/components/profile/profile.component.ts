import { Component } from '@angular/core';
import { User } from 'src/app/interfaces/user.interface';
import { LocalStorageService } from 'src/app/services/local-storage.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent {
  currentUser: User | null;

  constructor(private localStorageService: LocalStorageService) {
    this.currentUser = this.localStorageService.getCurrentUser();
  }
}
