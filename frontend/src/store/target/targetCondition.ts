import { Individuals } from '../individuals/types'
import { IBasePoints, BasePoints } from '../basePoints/types'
import { Nature } from '../nature/types'
import { Item } from '../items/types'
import { Weather, Field } from '../weatherFields/types'

export class TargetCondition {
  readonly target: string;
  readonly level: number;
  readonly individuals: Individuals;
  readonly basePoints: IBasePoints;
  readonly ability: string;
  readonly nature: Nature;
  readonly item: Item;
  readonly condition: string;
  readonly weather: Weather;
  readonly field: Field;

  enable(): boolean {
    if (this.target.length == 0) return false;
    if (this.level == 0) return false;
    if (this.ability.length == 0) return false;
    if (!this.nature.enable()) return false;
    if (!this.item.enable()) return false;
    if (this.condition.length == 0) return false;
    if (!this.weather.enable()) return false;
    if (!this.field.enable()) return false;
    return true;
  }

  constructor(
      target: string,
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
      this.target = target;
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
  static default(): TargetCondition {
    return new TargetCondition(
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
}
