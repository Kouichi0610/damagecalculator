import axios from 'axios'

export class SpeciesLoader {
  private name: string;

  constructor(name: string) {
    this.name = name;
  }

  public load(): Promise<Species> {
    return new Promise((resolve, reject) => {
      axios.get('get_species', {
        params: {
          Name: this.name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let sp = JSON.parse(json);
        resolve(new Species(sp));
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(new Species(null));
      });
    });
  }
}

export class Species {
  private _name: string = '';
  private _types: string = '';
  private _weight: number = 0;
  private _species: number[] = [0,0,0,0,0,0];
  private _abilities: string[] = [];

  hasTarget(): boolean {
    return this._name.length > 0;
  }
  hp(): number {return this._species[0];}
  attack(): number {return this._species[1];}
  defense(): number {return this._species[2];}
  spAttack(): number {return this._species[3];}
  spDefense(): number {return this._species[4];}
  speed(): number {return this._species[5];}
  name(): string {return this._name;}
  types(): string {return this._types;}
  weight(): number {return this._weight;}
  abilities(): string[] {return this._abilities;}

  static default(): Species {
    return new Species(null);
  }
  constructor(data: any) {
    if (data == null) {
      return;
    }
    this._name = data.Name;
    this._types = data.Types;
    this._weight = data.Weight;
    this._species = data.Species;
    this._abilities = data.Abilities;
  }
}
