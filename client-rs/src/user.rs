struct User {
    id: String,
}

impl User {
    fn new<T: Into<String>>(id: T) -> Self {
        User { id: id.into() }
    }
}
