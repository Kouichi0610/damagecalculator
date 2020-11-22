export interface BasePointsState {
}

/*
  基礎ポイント
  HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさ
  それぞれに割り振ることができる
  各最大252まで
  合計508まで
*/
export class BasePoints implements IBasePointsLimit {
  hp(): BasePoint { return this.basePoints[0] }
  attack(): BasePoint { return this.basePoints[1] }
  defense(): BasePoint { return this.basePoints[2] }
  spAttack(): BasePoint { return this.basePoints[3] }
  spDefense(): BasePoint { return this.basePoints[4] }
  speed(): BasePoint { return this.basePoints[5] }

  set(index: number, value: number) {
    let next = Math.min(value, this.basePoints[index].limit());
    this.basePoints[index] = new BasePoint(next, this);
  }

  private basePoints: BasePoint[];
  readonly max: number = 508;

  constructor() {
    this.basePoints = [];
    for (var i = 0; i < 6; i++) {
      this.basePoints.push(new BasePoint(0, this));
    }
  }

  limit(without: BasePoint): number {
    var total = 0;
    for (var i = 0; i < this.basePoints.length; i++) {
      if (without == this.basePoints[i]) continue;
      total += this.basePoints[i].value;
    }
    return Math.min((this.max - total), 252);
  }
}

interface IBasePointsLimit {
  limit(without: BasePoint): number;
}

export class BasePoint {
  readonly max: number = 252;
  readonly value: number = 0;
  private _limit: IBasePointsLimit;

  limit(): number {
    return this._limit.limit(this);
  }

  constructor(value: number, limit: IBasePointsLimit) {
    this.value = value;
    this._limit = limit;
    if (this.value > this.max) this.value = this.max;
    if (this.value < 0) this.value = 0;
  }
}

