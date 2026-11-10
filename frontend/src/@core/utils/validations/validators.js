export const validatorIsImageFile = file => {
  if (!file?.type) {
    return false
  }
  const validTypes = ["image/jpeg", "image/png", "image/webp"];
  return validTypes.includes(file.type);
};