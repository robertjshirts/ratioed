export function useRatio(likeCount: Ref<number>, dislikeCount: Ref<number>) {
  // Ratio posts if the like count is less than the dislike count
  const ratioed = ref(isRatioed(likeCount.value, dislikeCount.value));

  // Update the ratio whenever the like or dislike count changes
  watch([likeCount, dislikeCount], () => {
    ratioed.value = isRatioed(likeCount.value, dislikeCount.value);
  });

  return ratioed;
}

function isRatioed(likeCount: number, dislikeCount: number) {
  const total = likeCount + dislikeCount;
  if (total < 5) {
    return false;
  }
  let ratio = likeCount / dislikeCount;
  return ratio < 1;
}
