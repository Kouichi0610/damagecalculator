/*
  対象ポケモンに関するStore
  TODO:store主体だと手続き型になるので機能クラスにアクセスする形式に変更
  TODO:target/store改修
  TODO:action.vueからアクセス
  TODO:moveResult.vueに現在の技(+index)を保管
  TODO:defenderDamagesをこっち側に
*/
export interface TargetState {
  name: string;
  types: string[];
  species: Species;
  weight: number;
  abilities: string[];
  currentAbility: string;

  nature: string;
  individuals: Individuals;
  basePoints: BasePoints;

  // 基礎ポイント0～252(4刻み)
  statePatterns: StatePatterns;

  //moves: string[];
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

export class StatePatterns {
  private h: StatePattern;
  private at: StatePattern;
  private df: StatePattern;
  private sa: StatePattern;
  private sd: StatePattern;
  private sp: StatePattern;

  hp(basePoint: number): number {
    return this.h.value(basePoint);
  }
  attack(basePoint: number): number {
    return this.at.value(basePoint);
  }
  defense(basePoint: number): number {
    return this.df.value(basePoint);
  }
  spAttack(basePoint: number): number {
    return this.sa.value(basePoint);
  }
  spDefense(basePoint: number): number {
    return this.sd.value(basePoint);
  }
  speed(basePoint: number): number {
    return this.sp.value(basePoint);
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

// レベル、名前、性格、個体値から能力値のパターンを取得
// TODO:こっちにaxios持たせる
export class StatsPatternArgs {
  level: number;
  name: string;
  nature: string;
  individual: string;
  constructor(level: number, name: string, nature: string, hp: number, at: number, df: number, sa: number, sd: number, sp: number) {
    this.level = level;
    this.name = name;
    this.nature = nature;
    this.individual = "Max";
    if (at == 0) {
      this.individual = "Weakest";
    }
    if (sp == 0) {
      this.individual = "Slowest";
    }
  }
}

export interface Species {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
export interface Individuals {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
export interface BasePoints {
  hp: number;
  at: number;
  df: number;
  sa: number;
  sd: number;
  sp: number;
}
