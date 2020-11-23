export interface ISpecies {
  readonly hp: number;
  readonly attack: number;
  readonly defense: number;
  readonly spAttack: number;
  readonly spDefense: number;
  readonly speed: number;
}

export class Species implements ISpecies {
  readonly hp: number = 0;
  readonly attack: number = 0;
  readonly defense: number = 0;
  readonly spAttack: number = 0;
  readonly spDefense: number = 0;
  readonly speed: number = 0;

  constructor(hp: number, attack: number, defense: number, spAttack: number, spDefense: number, speed: number) {
    this.hp = hp;
    this.attack = attack;
    this.defense = defense;
    this.spAttack = spAttack;
    this.spDefense = spDefense;
    this.speed = speed;
  }
}
