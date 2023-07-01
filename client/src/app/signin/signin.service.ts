import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class SigninService {
  constructor(private http: HttpClient) {}

  signin(provider: string, idToken: string, accessToken: string) {
    return this.http.post(
      `${environment.API_URL}/auth/signin`,
      {},
      {
        headers: {
          'X-Provider': provider,
          'X-IdToken': idToken,
          'X-AccessToken': accessToken,
        },
        observe: 'response',
      }
    );
  }
}
