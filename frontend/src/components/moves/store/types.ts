
export class MoveInfo {
  name: string;
  type: string;
  power: number;
  accuracy: number;
  category: string;
  mention: string;

  constructor(data: any) {
    this.name = data.name;
    this.type = data.type;
    this.power = data.power;
    this.accuracy = data.accuracy;
    this.category = data.category;
    this.mention = data.mention;
  }

  ToString(): string {
    let res = '';
    res += this.name + ' ' + this.type + ' ' + this.power + ' 命中' + this.accuracy + ' ' + this.category + ' ' + this.mention;
    return res;
  }
}

export interface MovesState {
  // わざ一覧
}

