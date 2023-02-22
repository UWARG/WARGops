import { Profile } from '@/types';
export interface Account {
  name: "WEEF";
  id: string;
  source: string;
  active: true;
  allocation_date: Date;
  expiry_ate: Date;
  creator: string;
  point_of_contact: string; // must be email
  waterloo_id: string;
}

export interface NewTransaction {
  id: string;
  name: string;
  account_id: string;
  amount: number;
  type: number;
  status: number;
  notes: string;
}


export interface Transaction {
  id: string;
  name: string;
  amount: number;
  approval_date: Date;
  creation_date: Date;
  notes: string;
  payment_date: Date;
  rejected_date: Date;
  status: number;
  type: number;
}

export interface Profile {
  RawData: {
    accent_color: number;
    avatar: string;
    avatar_decoration: any;
    banner: any;
    banner_color: string;
    discriminator: string;
    display_name: any;
    flags: number;
    id: string;
    locale: string;
    mfa_enabled: boolean;
    premium_type: number;
    public_flags: number;
    roles: string[];
    username: string;
  };
  Provider: string;
  Email: string;
  Name: string;
  FirstName: string;
  LastName: string;
  NickName: string;
  Description: string;
  UserID: string;
  AvatarURL: string;
  Location: string;
  AccessToken: string;
  AccessTokenSecret: string;
  RefreshToken: string;
  ExpiresAt: string;
  IDToken: string;
}

