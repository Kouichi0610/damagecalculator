import { INature } from './nature';

export class StatsCalculator {
  constructor() {
  }
  Calculate(nature: INature, level: number, species: number[], individuals: number[], basePoints: number[]): number[] {
    let hp = this.hp(level, species[0], individuals[0], basePoints[0]);
    let at = nature.Attack(this.stats(level, species[1], individuals[1], basePoints[1]));
    let df = nature.Defense(this.stats(level, species[2], individuals[2], basePoints[2]));
    let sa = nature.SpAttack(this.stats(level, species[3], individuals[3], basePoints[3]));
    let sd = nature.SpDefense(this.stats(level, species[4], individuals[4], basePoints[4]));
    let sp = nature.Speed(this.stats(level, species[5], individuals[5], basePoints[5]));

    let res: number[] = [hp, at, df, sa, sd, sp];
    return res;
  }

  private hp(l: number, s: number, i: number, b: number): number {
    let x = s * 2 + i + b / 4;
    let y = x * l / 100.0 + l + 10;
    return Math.floor(y);
  }
  private stats(l: number, s: number, i: number, b: number): number {
    let x = s * 2 + i + b / 4;
    let y = Math.floor(x * l / 100.0 + 5);
    return Math.floor(y);
  }
}
