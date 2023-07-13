import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { SigninResponse } from '../interfaces/signin-response.interface';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class SigninService {
  constructor(private http: HttpClient) {}

  signin(
    provider: string,
    idToken: string,
    accessToken: string
  ): Observable<SigninResponse> {
    return this.http.post<SigninResponse>(
      `${environment.API_URL}/auth/signin`,
      {},
      {
        headers: {
          'X-Provider': provider,
          'X-IdToken': idToken,
          'X-AccessToken': accessToken,
        },
      }
    );
  }
}
