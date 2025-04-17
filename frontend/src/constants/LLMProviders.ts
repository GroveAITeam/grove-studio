// LLM提供商相关常量信息

export interface ProviderInfo {
  id: string;
  name: string;
  icon: string;
  description: string;
  endpoint: string;
  models: string[];
}

// LLM提供商列表
export const LLM_PROVIDERS: ProviderInfo[] = [
  {
    id: 'openai',
    name: 'OpenAI',
    icon: '/src/assets/images/providers/openai.png',
    description: '使用OpenAI API访问GPT-3.5, GPT-4等模型。',
    endpoint: 'https://api.openai.com/v1',
    models: ['gpt-4', 'gpt-3.5-turbo']
  },
  {
    id: 'deepseek',
    name: 'DeepSeek',
    icon: '/src/assets/images/providers/deepseek.png',
    description: '使用DeepSeek API访问DeepSeek系列模型。',
    endpoint: 'https://api.deepseek.com/v1',
    models: ['deepseek-chat', 'deepseek-coder']
  }
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
