export interface IDamageResult {
  readonly target: string;
  readonly damageMin: number;
  readonly damageMax: number;
  readonly rateMin: number;
  readonly rateMax: number;
  readonly determineCount: number;
  toString(): string;
}