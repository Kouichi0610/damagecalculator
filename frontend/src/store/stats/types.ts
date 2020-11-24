/*
  TODO:stats取得クラス

  storeは結果保存のみ
*/
import axios from 'axios'
import { Nature } from '../nature/types'
import { Individuals } from '../individuals/types'
import { IBasePoints } from '../basePoints/types'

export class StatsPatterns {
  statsArray(basePoints: IBasePoints): number[] {
    return [
      this.h.value(basePoints.hp()),
      this.at.value(basePoints.attack()),
      this.df.value(basePoints.defense()),
      this.sa.value(basePoints.spAttack()),
      this.sd.value(basePoints.spDefense()),
      this.sp.value(basePoints.speed()),
    ];
  }

  private h: StatePattern;
  private at: StatePattern;
  private df: StatePattern;
  private sa: StatePattern;
  private sd: StatePattern;
  private sp: StatePattern;

  constructor(data: any) {
    this.h = new StatePattern(data.hp);
    this.at = new StatePattern(data.attack);
    this.df = new StatePattern(data.defense);
    this.sa = new StatePattern(data.sp_attack);
    this.sd = new StatePattern(data.sp_defense);
    this.sp = new StatePattern(data.speed);
  }

  static default(): StatsPatterns {
    return new StatsPatterns({
      hp: [],
      attack: [],
      defense: [],
      sp_attack: [],
      sp_defense: [],
      speed: [],
    });
  }
}
class StatePattern {
  private values: number[];
  value(basePoint: number): number {
    if (this.values.length == 0) {
      return 0;
    }
    let idx = basePoint/4;
    return this.values[idx];
  }
  constructor(values: number[]) {
    this.values = values;
  }
}

export class StatsPatternsLoader {
  private level: number;
  private name: string;
  private individuals: Individuals;
  private nature: Nature;

  public enable(): boolean {
    if (this.name.length == 0) return false;
    if (this.nature.name.length == 0) return false;
    if (this.individuals.type().length == 0) return false;
    return true;
  }

  constructor(level: number, name: string, individuals: Individuals, nature: Nature) {
    this.level = level;
    this.name = name;
    this.individuals = individuals;
    this.nature = nature;
  }
  public load(): Promise<StatsPatterns> {
    return new Promise((resolve, reject) => {
      axios.get('stats_pattern', {
        params: {
          Level: this.level,
          Name: this.name,
          Nature: this.nature.name,
          Individual: this.individuals.type(),
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let res = JSON.parse(json);
        resolve(new StatsPatterns(res));
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(StatsPatterns.default());
      });
    });
  }
}
