export interface DamageRateInfo {
  readonly DamageMin: number;
  readonly DamageMax: number;
  readonly RateMin: number;
  readonly RateMax: number;
  readonly DetermineCount: number;

  toString(): string;
}