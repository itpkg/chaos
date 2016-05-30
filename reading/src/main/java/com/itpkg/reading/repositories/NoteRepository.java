package com.itpkg.reading.repositories;

import com.itpkg.core.models.Locale;
import com.itpkg.reading.models.Note;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by flamen on 16-5-30.
 */
public interface NoteRepository extends CrudRepository<Note, Long> {
}
