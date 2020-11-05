/*
  ランク補正
  TODO:サーバー側で賄いたい
*/
export class Rank {
  value: number;
  constructor(value: number) {
    this.value = value;
  }

  RankedStatus(status: number): number {
    if (this.value < 0) {
      let x = -this.value + 2;
      let res = status * 2 / x;
      return Math.floor(res);
    }
    {
      let x = this.value + 2;
      let res = status * x / 2;
      return Math.floor(res);
    }
  }
}

