import { Rank } from '../../../rank/rank'

export interface SpeedInfo {
  info: string;
  speed: number;
}

export interface SpeedOrderState {
  targetSpeed: number;
  list: SpeedInfo[];
  // とくせい自身 とくせい相手 もちもの[]
  abilityOwner: SpeedCorrector;
  abilityOther: SpeedCorrector;
  item: SpeedCorrector[];
}

export function DefaultSpeedCorrector(): SpeedCorrector {
  return new SpeedCorrector('', 0, 1.0);
}

export class SpeedCorrector {
  comment: string;
  rank: Rank;
  magnification: number;
  constructor(comment: string, rank: number, magnification: number) {
    this.comment = comment;
    this.rank = new Rank(rank);
    this.magnification = magnification;
  }
  correct(stats: number): number {
    let res = this.rank.RankedStatus(stats);
    return Math.floor(res * this.magnification);
  }
  toString(): string {
    return '' + this.comment + ' Rank:' + this.rank.value + ' Mag:' + this.magnification;
  }
}
