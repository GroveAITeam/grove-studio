export namespace config {

  export class AppConfig {
    version: string
    data_path: string
    debug_mode: boolean

    static createFrom(source: any = {}) {
      return new AppConfig(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.version = source.version
      this.data_path = source.data_path
      this.debug_mode = source.debug_mode
    }
  }

}

export namespace models {

  export class CloudLLMModel {
    id: number
    // Go type: time
    created_at: any
    // Go type: time
    updated_at: any
    name: string
    provider: string
    endpoint: string
    api_key: string
    enabled: boolean

    static createFrom(source: any = {}) {
      return new CloudLLMModel(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.id = source.id
      this.created_at = this.convertValues(source.created_at, null)
      this.updated_at = this.convertValues(source.updated_at, null)
      this.name = source.name
      this.provider = source.provider
      this.endpoint = source.endpoint
      this.api_key = source.api_key
      this.enabled = source.enabled
    }

    convertValues(a: any, classs: any, asMap: boolean = false): any {
      if (!a) {
        return a
      }
      if (a.slice && a.map) {
        return (a as any[]).map(elem => this.convertValues(elem, classs))
      }
      else if (typeof a === 'object') {
        if (asMap) {
          for (const key of Object.keys(a)) {
            a[key] = new classs(a[key])
          }
          return a
        }
        return new classs(a)
      }
      return a
    }
  }
  export class Conversation {
    id: number
    // Go type: time
    created_at: any
    // Go type: time
    updated_at: any
    title: string

    static createFrom(source: any = {}) {
      return new Conversation(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.id = source.id
      this.created_at = this.convertValues(source.created_at, null)
      this.updated_at = this.convertValues(source.updated_at, null)
      this.title = source.title
    }

    convertValues(a: any, classs: any, asMap: boolean = false): any {
      if (!a) {
        return a
      }
      if (a.slice && a.map) {
        return (a as any[]).map(elem => this.convertValues(elem, classs))
      }
      else if (typeof a === 'object') {
        if (asMap) {
          for (const key of Object.keys(a)) {
            a[key] = new classs(a[key])
          }
          return a
        }
        return new classs(a)
      }
      return a
    }
  }
  export class Message {
    id: number
    // Go type: time
    created_at: any
    // Go type: time
    updated_at: any
    conversation_id: number
    role: string
    content: string

    static createFrom(source: any = {}) {
      return new Message(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.id = source.id
      this.created_at = this.convertValues(source.created_at, null)
      this.updated_at = this.convertValues(source.updated_at, null)
      this.conversation_id = source.conversation_id
      this.role = source.role
      this.content = source.content
    }

    convertValues(a: any, classs: any, asMap: boolean = false): any {
      if (!a) {
        return a
      }
      if (a.slice && a.map) {
        return (a as any[]).map(elem => this.convertValues(elem, classs))
      }
      else if (typeof a === 'object') {
        if (asMap) {
          for (const key of Object.keys(a)) {
            a[key] = new classs(a[key])
          }
          return a
        }
        return new classs(a)
      }
      return a
    }
  }

}

export namespace services {

  export class CloudLLMModelPageResult {
    total: number
    items: models.CloudLLMModel[]

    static createFrom(source: any = {}) {
      return new CloudLLMModelPageResult(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.total = source.total
      this.items = this.convertValues(source.items, models.CloudLLMModel)
    }

    convertValues(a: any, classs: any, asMap: boolean = false): any {
      if (!a) {
        return a
      }
      if (a.slice && a.map) {
        return (a as any[]).map(elem => this.convertValues(elem, classs))
      }
      else if (typeof a === 'object') {
        if (asMap) {
          for (const key of Object.keys(a)) {
            a[key] = new classs(a[key])
          }
          return a
        }
        return new classs(a)
      }
      return a
    }
  }
  export class ConversationPageResult {
    total: number
    items: models.Conversation[]

    static createFrom(source: any = {}) {
      return new ConversationPageResult(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.total = source.total
      this.items = this.convertValues(source.items, models.Conversation)
    }

    convertValues(a: any, classs: any, asMap: boolean = false): any {
      if (!a) {
        return a
      }
      if (a.slice && a.map) {
        return (a as any[]).map(elem => this.convertValues(elem, classs))
      }
      else if (typeof a === 'object') {
        if (asMap) {
          for (const key of Object.keys(a)) {
            a[key] = new classs(a[key])
          }
          return a
        }
        return new classs(a)
      }
      return a
    }
  }
  export class MessagePageResult {
    items: models.Message[]

    static createFrom(source: any = {}) {
      return new MessagePageResult(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.items = this.convertValues(source.items, models.Message)
    }

    convertValues(a: any, classs: any, asMap: boolean = false): any {
      if (!a) {
        return a
      }
      if (a.slice && a.map) {
        return (a as any[]).map(elem => this.convertValues(elem, classs))
      }
      else if (typeof a === 'object') {
        if (asMap) {
          for (const key of Object.keys(a)) {
            a[key] = new classs(a[key])
          }
          return a
        }
        return new classs(a)
      }
      return a
    }
  }
  export class MessageRequestParams {
    cloud_llm_id: number
    conversation_id: number
    question: string
    model_name: string
    temperature: number
    max_completion_tokens: number
    history_length: number

    static createFrom(source: any = {}) {
      return new MessageRequestParams(source)
    }

    constructor(source: any = {}) {
      if (typeof source === 'string') { source = JSON.parse(source) }
      this.cloud_llm_id = source.cloud_llm_id
      this.conversation_id = source.conversation_id
      this.question = source.question
      this.model_name = source.model_name
      this.temperature = source.temperature
      this.max_completion_tokens = source.max_completion_tokens
      this.history_length = source.history_length
    }
  }

}
