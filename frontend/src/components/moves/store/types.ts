export class MoveInfos {
  infos: MoveInfo[];
  constructor(infos: MoveInfo[]) {
    this.infos = infos;
    this.infos.sort(function(a, b) {
      if (a.power > b.power) return -1;
      if (a.power < b.power) return 1;
      return 0;
    })
  }

  physicals(): MoveInfo[] {
    let res: MoveInfo[] = [];
    for (var i = 0; i < this.infos.length; i++) {
      let category = this.infos[i].category;
      if (category == 'ぶつり') {
        res.push(this.infos[i])
      }
    }
    return res;
  }
  specials(): MoveInfo[] {
    let res: MoveInfo[] = [];
    for (var i = 0; i < this.infos.length; i++) {
      let category = this.infos[i].category;
      if (category != 'ぶつり') {
        res.push(this.infos[i])
      }
    }
    return res;
  }
}

export class MoveInfo {
  name: string;
  type: string;
  power: number;
  accuracy: number;
  category: string;
  mention: string;

  static empty(): MoveInfo {
    return new MoveInfo({name: '', type: '', power: 0, accurcy: 0, category: '', mention: ''});
  }

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
  info: MoveInfos;
}

