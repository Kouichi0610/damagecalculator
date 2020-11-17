import axios from 'axios'
import { Individuals } from './individuals'
import { BasePoints } from './basePoints'

export class StatePatternsLoader {
  private level: number;
  private name: string;
  private nature: string;
  private individuals: string;

  constructor(level: number, name: string, nature: string, individuals: Individuals) {
    this.level = level;
    this.name = name;
    this.nature = nature;
    this.individuals = individuals.type();
  }

  public load(): Promise<StatePatterns> {
    return new Promise((resolve, reject) => {
      axios.get('stats_pattern', {
        params: {
          Level: this.level,
          Name: this.name,
          Nature: this.nature,
          Individual: this.individuals,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let res = JSON.parse(json);
        resolve(new StatePatterns(res));
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(StatePatterns.default());
      });
    });
  }
}

/*
  基礎ポイントごとの能力値一覧
*/
export class StatePatterns {
  private h: StatePattern;
  private at: StatePattern;
  private df: StatePattern;
  private sa: StatePattern;
  private sd: StatePattern;
  private sp: StatePattern;

  hp(basePoint: BasePoints): number {
    return this.h.value(basePoint.hp);
  }
  attack(basePoint: BasePoints): number {
    return this.at.value(basePoint.attack);
  }
  defense(basePoint: BasePoints): number {
    return this.df.value(basePoint.defense);
  }
  spAttack(basePoint: BasePoints): number {
    return this.sa.value(basePoint.spAttack);
  }
  spDefense(basePoint: BasePoints): number {
    return this.sd.value(basePoint.spDefense);
  }
  speed(basePoint: BasePoints): number {
    return this.sp.value(basePoint.speed);
  }

  static default(): StatePatterns {
    return new StatePatterns({
      hp: [],
      attack: [],
      defense: [],
      sp_attack: [],
      sp_defense: [],
      speed: [],
    });
  }
  constructor(data: any) {
    this.h = new StatePattern(data.hp);
    this.at = new StatePattern(data.attack);
    this.df = new StatePattern(data.defense);
    this.sa = new StatePattern(data.sp_attack);
    this.sd = new StatePattern(data.sp_defense);
    this.sp = new StatePattern(data.speed);
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
