export interface IndividualsState {
  target: string;
  targetsIndividuals: Individuals;
}

export class Individuals {
  readonly hp: number = 31;
  readonly attack: number = 31;
  readonly defense: number = 31;
  readonly spAttack: number = 31;
  readonly spDefense: number = 31;
  readonly speed: number = 31;

  private constructor(hp: number, at: number, df: number, sa: number, sd: number, sp: number) {
    this.hp = hp;
    this.attack = at;
    this.defense = df;
    this.spAttack = sa;
    this.spDefense = sd;
    this.speed = sp;
  }
  static default(): Individuals {
    return new Individuals(31,31,31,31,31,31);
  }
  changeSlowest(): Individuals {
    return new Individuals(
      this.hp,
      this.attack,
      this.defense,
      this.spAttack,
      this.spDefense,
      this.isSlowest() ? 31 : 0
    );
  }
  changeWeakest(): Individuals {
    return new Individuals(
      this.hp,
      this.isWeakest() ? 31: 0,
      this.defense,
      this.spAttack,
      this.spDefense,
      this.speed
    );
  }

  isSlowest(): boolean {
    return this.speed == 0;
  }
  isWeakest(): boolean {
    return this.attack == 0;
  }
  public type(): string {
    if (this.isWeakest() && this.isSlowest()) {
      return "WeakestSlowest";
    }
    if (this.isWeakest()) {
      return "Weakest";
    }
    if (this.isSlowest()) {
      return "Slowest";
    }
    return "Max";
  }
  public toString(): string {
    return '' + this.hp + '-' + this.attack + '-' + this.defense + '-' + this.spAttack + '-' + this.spDefense + '-' + this.speed;
  }
}
