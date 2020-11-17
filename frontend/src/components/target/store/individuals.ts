/*
  個体値
*/
export class Individuals {
  private _hp: number = 31;
  private _at: number = 31;
  private _df: number = 31;
  private _sa: number = 31;
  private _sd: number = 31;
  private _sp: number = 31;

  public weakdef() {
    this._df = 0;
  }

  public isSlowest(): boolean {
    return this._sp == 0;
  }
  public isWeakest(): boolean {
    return this._at == 0;
  }
  public slowest(isSlowest: boolean) {
    this._sp = isSlowest ? 0 : 31;
  }
  public weakest(isWeakest: boolean) {
    this._at = isWeakest ? 0 : 31;
  }

  public type(): string {
    if (this._at == 0 && this._sp == 0) {
      return "WeakestSlowest";
    }
    if (this._at == 0) {
      return "Weakest";
    }
    if (this._sp == 0) {
      return "Slowest";
    }
    return "Max";
  }
  public hp(): number {return this._hp;}
  public attack(): number {return this._at;}
  public defense(): number {return this._df;}
  public spAttack(): number {return this._sa;}
  public spDefense(): number {return this._sd;}
  public speed(): number {return this._sp;}

  public toString(): string {
    return '' + this.hp() + '-' + this.attack() + '-' + this.defense() + '-' + this.spAttack() + '-' + this.spDefense() + '-' + this.speed();
  }

  constructor() {
  }
}