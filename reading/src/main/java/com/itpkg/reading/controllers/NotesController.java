package com.itpkg.reading.controllers;

import com.itpkg.reading.models.Note;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * Created by flamen on 16-5-30.
 */
@RestController
@RequestMapping(value = "/reading")
public class NotesController {
    @RequestMapping(value = "/notes", method = RequestMethod.GET)
    public List<Note> index(){
        //todo
        List<Note> notes = new ArrayList<>();
        for(long i=0; i<10; i++){
            Note n = new Note();
            n.setUpdatedAt(new Date());
            n.setId(i);
            n.setTitle(String.format("note title %d", i));
            n.setTitle(String.format("note body %d", i));
            notes.add(n);
        }
        return notes;
    }
}
