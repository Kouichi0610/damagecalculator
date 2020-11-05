/*
  対象となるポケモンを選択するためのStore
*/
export interface TargetSelectState {
  initialTotal: number; // 種族値合計(この値以下をリストから除く)
  initialType: string;  // タイプ(このタイプのみリストを作成する)
  candidates: Species[]; // 候補者一覧
}

// 候補者を絞り込むためのフィルター
export class CandidatesFilter {
  type: string = 'すべて';  // タイプ条件を満たすもの
  total: number = 500;     // 種族値合計を満たすもの

  constructor(type: string, total: number) {
    this.type = type;
    this.total = total;
  }
}
// 候補者情報
export class Species {
  name: string;
  types: string;
  hp: number;
  attack: number;
  defense: number;
  spAttack: number;
  spDefense: number;
  speed: number;
  constructor (d: any) {
    this.name = d.name;
    this.types = d.types;
    this.hp = d.hp;
    this.attack = d.attack;
    this.defense = d.defense;
    this.spAttack = d.sp_attack;
    this.spDefense = d.sp_defense;
    this.speed = d.speed;
  }
  toString(): string {
    return this.name + ' ' + this.types
    + ' HP:' + this.hp
    + ' AT:' + this.attack
    + ' DF:' + this.defense
    + ' SA:' + this.spAttack
    + ' SD:' + this.spDefense
    + ' SP:' + this.speed;
  }
}
