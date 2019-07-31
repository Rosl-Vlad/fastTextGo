#include <iostream>
#include <istream>
#include "fasttext.h"
#include "real.h"
#include <streambuf>
#include <cstring>
#include <map>

extern "C" {

std::map<std::string, fasttext::FastText*> g_fasttext_model;

void load_model(char *name, char *path) {
	fasttext::FastText *model=new fasttext::FastText();
	model->loadModel(std::string(path));
	g_fasttext_model[std::string(name)]=model;
}

int get_sentence_vector(char* name, char* sentence, float* buf) {
	try {
		g_fasttext_model.at(std::string(name))->sentenceVectors(std::string(sentence), buf);
		return 0;
	} catch (const std::exception& e) {
		return 1;
	}
	return 1;
}
}
