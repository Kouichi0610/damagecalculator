export interface BasePointsState {
  basePoints: BasePoints
}

const tags: string[] = [
  'HP',
  '攻撃',
  '防御',
  '特防',
  '特防',
  '素早さ',
]

// TODO: implements readonlyBasePoints

/*
  基礎ポイント
  HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさ
  それぞれに割り振ることができる
  各最大252まで
  合計508まで
*/
export class BasePoints implements IBasePointsLimit {
  hp(): IBasePoint { return this.basePoints[0] }
  attack(): IBasePoint { return this.basePoints[1] }
  defense(): IBasePoint { return this.basePoints[2] }
  spAttack(): IBasePoint { return this.basePoints[3] }
  spDefense(): IBasePoint { return this.basePoints[4] }
  speed(): IBasePoint { return this.basePoints[5] }
  array(): IBasePoint[] { return this.basePoints; }

  total(): number {
    var res = 0;
    for (var i = 0; i < this.basePoints.length; i++) {
      res += this.basePoints[i].value;
    }
    return res;
  }

  set(index: number, value: number) {
    let next = Math.min(value, this.basePoints[index].limit());
    this.basePoints[index].value = next;
  }

  private basePoints: BasePoint[];
  readonly max: number = 508;

  constructor() {
    this.basePoints = [];
    for (var i = 0; i < 6; i++) {
      this.basePoints.push(new BasePoint(0, i, tags[i], this));
    }
  }

  limit(without: BasePoint): number {
    var total = 0;
    for (var i = 0; i < this.basePoints.length; i++) {
      if (without.index == this.basePoints[i].index) continue;
      total += this.basePoints[i].value;
    }
    return Math.min((this.max - total), 252);
  }

  toString(): string {
    var res = '';
    for (var i = 0; i < this.basePoints.length; i++) {
      res += this.basePoints[i].tag + ':' + this.basePoints[i].value + ' ';
    }
    return res;
  }
}

interface IBasePointsLimit {
  limit(without: BasePoint): number;
}

export interface IBasePoint {
  readonly index: number;
  readonly tag: string;
  readonly value: number;
  limit(): number;
}

class BasePoint implements IBasePoint {
  readonly index: number;
  readonly tag: string;
  value: number = 0;
  readonly max: number = 252;
  private _limit: IBasePointsLimit;

  limit(): number {
    return this._limit.limit(this);
  }

  constructor(value: number, index: number, tag: string, limit: IBasePointsLimit) {
    this.value = value;
    this.index = index;
    this.tag = tag;
    this._limit = limit;
    if (this.value > this.max) this.value = this.max;
    if (this.value < 0) this.value = 0;
  }
}

