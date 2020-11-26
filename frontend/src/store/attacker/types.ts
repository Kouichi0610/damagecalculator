import axios from 'axios'

/*
  攻撃調整
  わざ4つまで
*/
export interface AttackerState {
  currentIndex: number;
  // TODO:配列では上手くいかない?
  moveA: string;
  moveB: string;
  moveC: string;
  moveD: string;

  attacker: string;
  physicals: Move[];
  specials: Move[];
}

export class Move {
  readonly name: string;
  readonly type: string;
  readonly power: number;
  readonly accuracy: number;
  readonly category: string;
  readonly mention: string;

  isPhysical(): boolean {
    return this.category == 'ぶつり';
  }
  isSpecial(): boolean {
    return this.category == 'とくしゅ';
  }

  constructor(name: string, type: string, power: number, accuracy: number, category: string, mention: string) {
    this.name = name;
    this.type = type;
    this.power = power;
    this.accuracy = accuracy;
    this.category = category;
    this.mention = mention;
  }

  toString(): string {
    return '' + this.name + ' ' + this.type + ' ' + this.category + ' 威力:' + this.power;
  }
}

export class MoveLoader {
  load(name: string): Promise<Move[]> {
    return new Promise((resolve, reject) => {
      axios.get('moves', {
        params: {
          Name: name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let datas = JSON.parse(json);
        let moves: Move[] = [];
        for (var i = 0; i < datas.length; i++) {
          moves.push(this.newMove(datas[i]));
        }
        resolve(moves);
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject([]);
      });
    });
  }

  private newMove(data: any) {
    return new Move(
      data.name,
      data.type,
      data.power,
      data.accuracy,
      data.category,
      data.mention
    );
  }

}
