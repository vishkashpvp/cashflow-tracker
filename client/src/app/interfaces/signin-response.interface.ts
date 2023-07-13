import { User } from './user.interface';

export interface SigninResponse {
  status: number;
  token: string;
  user: User;
}
