export class BasePoints {
  readonly hp:number = 0;
  readonly attack:number = 0;
  readonly defense:number = 0;
  readonly spAttack:number = 0;
  readonly spDefense:number = 0;
  readonly speed:number = 0;

  constructor(values: number[]) {
    this.hp = values[0];
    this.attack = values[1];
    this.defense = values[2];
    this.spAttack = values[3];
    this.spDefense = values[4];
    this.speed = values[5];
  }

  static default(): BasePoints {
    return new BasePoints([0,0,0,0,0,0])
  }
}