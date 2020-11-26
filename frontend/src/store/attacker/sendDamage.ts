import axios from 'axios'
import { Individuals } from '../individuals/types'
import { IBasePoints, BasePoints } from '../basePoints/types'
import { Nature } from '../nature/types'
import { Item } from '../items/types'
import { Weather, Field } from '../weatherFields/types'

export class SendDamages {
  readonly attacker: string;
  private level: number;
  private individuals: Individuals;
  private basePoints: IBasePoints;
  private ability: string;
  private nature: Nature;
  private item: Item;
  private condition: string;
  private weather: Weather;
  private field: Field;

  sendDamages(move: string): Promise<Result[]> {
    return new Promise((resolve, reject) => {
      axios.get('defender_damages', {
        params: {
          Level: this.level,
          BaseHP: this.basePoints.hp(),
          BaseAttack: this.basePoints.attack(),
          BaseDefense: this.basePoints.defense(),
          BaseSpAttack: this.basePoints.spAttack(),
          BaseSpDefense: this.basePoints.spDefense(),
          BaseSpeed: this.basePoints.speed(),
          Individuals: this.individuals.type(),
          Name: this.attacker,
          Move: move,
          Ability: this.ability,
          Nature: this.nature.name,
          Item: this.item.name,
          Condition: this.condition,
          Weather: this.weather.name,
          Field: this.field.name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let damages = JSON.parse(json);
        let res: Result[] = []

        for (var i = 0; i < damages.length; i++) {
          let d = new Result(damages[i]);
          res.push(d);
          //console.log('' + i + ' ' + d.toString());
        }
        resolve(res);
      })
      .catch((e) => {
        console.log('failed:' + e);
        let res: Result[] = []
        reject(res);
      });
    });
  }

  static default(): SendDamages {
    return new SendDamages(
      '',
      0,
      Individuals.default(),
      new BasePoints(),
      '',
      Nature.default(),
      Item.default(),
      '',
      Weather.default(),
      Field.default()
    )
  }

  constructor(
      attacker: string,
      level: number,
      individuals: Individuals,
      basePoints: IBasePoints,
      ability: string,
      nature: Nature,
      item: Item,
      condition: string,
      weather: Weather,
      field: Field,
    ) {
      this.attacker = attacker;
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
}

export class Result {
  readonly Target: string = '';
  readonly DamageMin: number = 0;
  readonly DamageMax: number = 0;
  readonly RateMin: number = 0;
  readonly RateMax: number = 0;
  readonly DetermineCount: number = 0;

  toString(): string {
    return this.Target
     + ' ダメージ' + this.DamageMin + '～' + this.DamageMax
     + ' ' + this.RateMin + '%～' + this.RateMax + '%'
     + ' 確定数' + this.DetermineCount;
  }

  constructor(d: any) {
    if (d == null) {
      return;
    }
    this.Target = d.target;
    this.DamageMin = d.damage_min;
    this.DamageMax = d.damage_max;
    this.RateMax = d.rate_max;
    this.RateMin = d.rate_min;
    this.DetermineCount = d.determine_count;
  }
}
