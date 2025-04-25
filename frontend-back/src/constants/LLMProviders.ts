// LLM提供商相关常量信息

export interface ProviderInfo {
  id: string;
  name: string;
  icon: string;
  endpoint: string;
  models: string[];
}

// LLM提供商列表
export const LLM_PROVIDERS: ProviderInfo[] = [
  {
    id: 'openai',
    name: 'OpenAI',
    icon: '/src/assets/images/providers/openai.png',
    endpoint: 'https://api.openai.com/v1',
    models: [
      'gpt-4o',
      'gpt-4o-mini',
      'gpt-4.1',
    ]
  },
  {
    id: 'deepseek',
    name: 'DeepSeek',
    icon: '/src/assets/images/providers/deepseek.png',
    endpoint: 'https://api.deepseek.com/v1',
    models: [
      'deepseek-chat',
      'deepseek-reasoner',
    ]
  },
  {
    id: 'siliconflow',
    name: '硅基流动(siliconflow)',
    icon: '/src/assets/images/providers/deepseek.png',
    endpoint: 'https://api.siliconflow.cn/v1',
    models: [
      'THUDM/GLM-Z1-32B-0414',
      'THUDM/GLM-4-32B-0414',
      'THUDM/GLM-Z1-Rumination-32B-0414',
      'THUDM/GLM-4-9B-0414',
      'THUDM/GLM-4-9B-0414',
      'Qwen/QwQ-32B',
      'Pro/deepseek-ai/DeepSeek-R1',
      'Pro/deepseek-ai/DeepSeek-V3',
      'deepseek-ai/DeepSeek-R1',
      'deepseek-ai/DeepSeek-V3',
      'deepseek-ai/DeepSeek-R1-Distill-Qwen-32B',
      'deepseek-ai/DeepSeek-R1-Distill-Qwen-14B',
      'deepseek-ai/DeepSeek-R1-Distill-Qwen-7B',
      'deepseek-ai/DeepSeek-R1-Distill-Qwen-1.5B',
      'Pro/deepseek-ai/DeepSeek-R1-Distill-Qwen-7B',
      'Pro/deepseek-ai/DeepSeek-R1-Distill-Qwen-1.5B',
    ],
  },
  {
    id: 'baidu',
    name: '百度',
    icon: '/src/assets/images/providers/baidu.png',
    endpoint: 'https://qianfan.baidubce.com/v2',
    models: [
      'ernie-x1-32k',
      'ernie-4.0-8k-latest',
      'ernie-4.0-8k',
      'ernie-4.0-turbo-8k-latest',
      'ernie-4.0-turbo-128k',
      'ernie-3.5-128k',
      'deepseek-v3',
      'deepseek-r1',
    ]
  },
  {
    id: 'tongyi',
    name: '通义千问',
    icon: '/src/assets/images/providers/tongyi.png',
    endpoint: 'https://dashscope.aliyuncs.com/compatible-mode/v1',
    models: [
      'qwen-max',
      'qwen-plus',
      'qwen-turbo',
      'qwen-long',
    ],
  },
  {
    id: 'hunyuan',
    name: '腾讯混元',
    icon: '/src/assets/images/providers/hunyuan.png',
    endpoint: 'https://api.hunyuan.cloud.tencent.com/v1',
    models: [
      'hunyuan-lite',
      'hunyuan-standard',
      'hunyuan-standard-256k',
      'hunyuan-turbo-latest',
      'hunyuan-large',
      'hunyuan-large-longcontext',
    ],
  },
  {
    id: 'huoshan',
    name: '火山引擎',
    icon: '/src/assets/images/providers/huoshan.png',
    endpoint: 'https://ark.cn-beijing.volces.com/api/v3/',
    models: [
      'doubao-1-5-pro-32k-250115',
      'doubao-1-5-pro-256k-250115',
      'doubao-1.5-lite-32k-250115',
      'deepseek-v3-250324',
      'deepseek-v3-241226',
    ],
  },
  {
    id: 'gemini',
    name: 'Google Gemini',
    icon: '/src/assets/images/providers/gemini.png',
    endpoint: 'https://generativelanguage.googleapis.com/v1beta/openai/',
    models: [
      'gemini-2.0-flash',
      'gemini-2.0-flash-lite',
      'gemini-1.5-flash',
      'gemini-1.5-pro',
    ],
  },
  {
    id: 'spark',
    name: '讯飞星火',
    icon: '/src/assets/images/providers/gemini.png',
    endpoint: 'https://spark-api-open.xf-yun.com/v1',
    models: [
      'x1',
      '4.0Ultra',
      'generalv3.5',
      'max-32k',
      'generalv3',
      'pro-128k',
      'lite',
    ],

  },
  {
    id: 'cohere',
    name: 'cohere',
    icon: '/src/assets/images/providers/cohere.png',
    endpoint: 'https://api.cohere.ai/compatibility/v1',
    models: [
      'command-r-plus',
      'command-r',
      'command',
      'command-light',
    ],
  },
  {
    id: 'kimi',
    name: 'Kimi',
    icon: '/src/assets/images/providers/kimi.png',
    endpoint: 'https://api.moonshot.cn/v1',
    models: [
      'kimi-latest-8k',
      'kimi-latest-32k',
      'kimi-latest-128k',
      'moonshot-v1-8k',
      'moonshot-v1-32k',
      'moonshot-v1-128k',
    ],
  },
  {
    id: 'baichuan',
    name: '百川智能',
    icon: '/src/assets/images/providers/baichuan.png',
    endpoint: 'https://api.baichuan-ai.com/v1',
    models: [
      'Baichuan4-Turbo',
      'Baichuan4-Air',
      'Baichuan4',
      'Baichuan3-Turbo',
      'Baichuan3-Turbo-128k',
      'Baichuan2-Turbo',
    ],
  },
];

// 根据提供商ID获取提供商信息
export function getProviderById(providerId: string): ProviderInfo | undefined {
  return LLM_PROVIDERS.find(provider => provider.id === providerId);
}

// 获取提供商图标
export function getProviderIcon(providerId: string): string {
  const provider = LLM_PROVIDERS.find(p => p.id === providerId);
  return provider ? provider.icon : '/src/assets/images/providers/default.png';
}

// 根据提供商ID获取可用模型列表
export function getProviderModels(providerId: string): string[] {
  const provider = getProviderById(providerId);
  return provider ? provider.models : [];
}
