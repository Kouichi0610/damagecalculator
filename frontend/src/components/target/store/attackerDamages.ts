import axios from 'axios';
import { BasePoints } from './basePoints'

/*
  耐久調整
  仮想敵からのダメージ一覧
*/
export class AttackerDamages {
  private defender: string;
  private level: number
  private individuals: string;
  private basePoints: BasePoints;
  private nature: string;
  private ability: string;
  private item: string;
  private condition: string;
  private weather: string;
  private field: string;
  constructor(defender: string, level: number, individuals: string, basePoints: BasePoints, nature: string, ability: string, item: string, condition: string, weather: string, field: string) {
    this.defender = defender;
    this.level = level;
    this.individuals = individuals;
    this.basePoints = basePoints;
    this.ability = ability;
    this.nature = nature;
    this.item = item;
    this.condition = condition;
    this.weather = weather;
    this.field = field;
  }

  attackerDamages(): Promise<AttackersResult[]> {
    return new Promise((resolve, reject) => {
      axios.get('attacker_damages', {
        params: {
          Level: this.level,
          BaseHP: this.basePoints.hp,
          BaseAttack: this.basePoints.attack,
          BaseDefense: this.basePoints.defense,
          BaseSpAttack: this.basePoints.spAttack,
          BaseSpDefense: this.basePoints.spDefense,
          BaseSpeed: this.basePoints.speed,
          Individuals: this.individuals,
          Name: this.defender,
          Ability: this.ability,
          Nature: this.nature,
          Item: this.item,
          Condition: this.condition,
          Weather: this.weather,
          Field: this.field,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let damages = JSON.parse(json);
        let res: AttackersResult[] = []

        for (var i = 0; i < damages.length; i++) {
          let d = new AttackersResult(damages[i]);
          res.push(d);
        }
        resolve(res);
      })
      .catch((e) => {
        console.log('error:' + e);
        let res : AttackersResult[] = [];
        reject(res);
      })
    });
  }
}

export class AttackersResult {
  readonly Target: string = '';
  readonly Move: string = '';
  readonly DamageMin: number = 0;
  readonly DamageMax: number = 0;
  readonly RateMin: number = 0;
  readonly RateMax: number = 0;
  readonly DetermineCount: number = 0;

  toString(): string {
    return this.Target
     + ' ' + this.Move
     + ' ダメージ' + this.DamageMin + '～' + this.DamageMax
     + ' ' + this.RateMin + '%～' + this.RateMax + '%'
     + ' 確定数' + this.DetermineCount;
  }

  constructor(d: any) {
    if (d == null) return;
    this.Target = d.target;
    this.Move = d.move;
    this.DamageMin = d.damage_min;
    this.DamageMax = d.damage_max;
    this.RateMax = d.rate_max;
    this.RateMin = d.rate_min;
    this.DetermineCount = d.determine_count;
  }

  Initialized(): boolean {
    return this.Target.length > 0;
  }

  static detault(): AttackersResult {
    return new AttackersResult(null);
  }
}

