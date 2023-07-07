import {
  FacebookLoginProvider,
  SocialAuthService,
  SocialUser,
} from '@abacritt/angularx-social-login';
import { Component, OnInit } from '@angular/core';
import { SigninService } from './signin.service';
import { LocalStorageService } from '../services/local-storage.service';

enum LoginProvider {
  GOOGLE = 'GOOGLE',
  FACEBOOK = 'FACEBOOK',
}

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
    private socialAuthService: SocialAuthService,
    private localStorageService: LocalStorageService
  ) {}

  ngOnInit() {
    this.socialAuthService.authState.subscribe((user: SocialUser) => {
      this.user = user;

      if (!user) this.loggedIn = false;

      switch (user.provider.toUpperCase()) {
        case LoginProvider.GOOGLE:
          this.signin(LoginProvider.GOOGLE, user.idToken, '');
          break;
        case LoginProvider.FACEBOOK:
          this.signin(LoginProvider.FACEBOOK, '', user.authToken);
          break;
        default:
          this.loggedIn = false;
          console.error('no such provider: ', user.provider);
          break;
      }
    });
  }

  signin(provider: string, idToken: string, accessToken: string) {
    this.service.signin(provider, idToken, accessToken).subscribe({
      next: (val: any) => {
        console.log('val :>> ', val);
        console.log('user :>> ', val.body.user);

        this.localStorageService.setAuthToken(val.body.token);

        this.loggedIn = true;
      },
      error: (err) => {
        console.log('err :>> ', err);
        console.log('msg :>> ', err.error.message);

        this.xError = true;
      },
    });
  }

  signInWithFB() {
    this.socialAuthService.signIn(FacebookLoginProvider.PROVIDER_ID);
  }
}
