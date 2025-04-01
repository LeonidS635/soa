package usecase

func (uc *PostsUseCase) DeletePost(postId, userId int32) error {
	return uc.storage.DeletePost(postId, userId)
}
