import { SocialAuthService, SocialUser } from '@abacritt/angularx-social-login';
import { Component, OnInit } from '@angular/core';
import { SigninService } from './signin.service';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.scss'],
})
export class SigninComponent implements OnInit {
  user!: SocialUser;
  loggedIn: boolean = false;
  xError: boolean = false;

  constructor(
    private service: SigninService,
    private socialAuthService: SocialAuthService
  ) {}

  ngOnInit() {
    this.socialAuthService.authState.subscribe((user: SocialUser) => {
      this.user = user;
      if (user == null) {
        this.loggedIn = false;
      } else {
        this.signin(user.idToken, user.provider);
      }
    });
  }

  signin(idToken: string, provider: string) {
    this.service.signin(idToken, provider).subscribe({
      next: (val: any) => {
        console.log('val :>> ', val);
        console.log('user :>> ', val.body.user);

        this.loggedIn = true;
      },
      error: (err) => {
        console.log('err :>> ', err);
        console.log('msg :>> ', err.error.message);

        this.xError = true;
      },
    });
  }
}
