// composables/useAutoResize.ts
import { ref } from 'vue';

export function useAutoResize() {
  const content = ref('');

  const autoResize = (event: Event) => {
    const textarea = event.target as HTMLTextAreaElement;
    textarea.style.height = 'auto'; // Reset the height to auto
    textarea.style.height = `${textarea.scrollHeight}px`; // Set the height to the scroll height
  };

  return {
    content,
    autoResize,
  };
}
